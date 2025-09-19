package streaming

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"job-board/backend/logger"
)

// VideoStreamer handles video streaming operations
type VideoStreamer struct {
	videoDirectory string
}

// NewVideoStreamer creates a new video streamer
func NewVideoStreamer(videoDirectory string) *VideoStreamer {
	return &VideoStreamer{
		videoDirectory: videoDirectory,
	}
}

// StreamVideo streams a video file with proper HTTP headers
func (vs *VideoStreamer) StreamVideo(w http.ResponseWriter, r *http.Request, videoID string) error {
	videoPath := filepath.Join(vs.videoDirectory, videoID+".mp4")

	// Check if video file exists
	fileInfo, err := os.Stat(videoPath)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Warn("Video file not found", "video_id", videoID, "path", videoPath)
			return fmt.Errorf("video not found")
		}
		logger.Error("Error checking video file", "video_id", videoID, "error", err)
		return fmt.Errorf("error accessing video file")
	}

	// Open the video file
	file, err := os.Open(videoPath)
	if err != nil {
		logger.Error("Error opening video file", "video_id", videoID, "error", err)
		return fmt.Errorf("error opening video file")
	}
	defer file.Close()

	// Set appropriate headers for video streaming
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Accept-Ranges", "bytes")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Handle range requests for video seeking
	rangeHeader := r.Header.Get("Range")
	if rangeHeader != "" {
		return vs.handleRangeRequest(w, r, file, fileInfo.Size(), rangeHeader)
	}

	// Stream the entire file
	logger.Info("Streaming video", "video_id", videoID, "size", fileInfo.Size())
	_, err = io.Copy(w, file)
	if err != nil {
		logger.Error("Error streaming video", "video_id", videoID, "error", err)
		return fmt.Errorf("error streaming video")
	}

	return nil
}

// handleRangeRequest handles HTTP range requests for video seeking
func (vs *VideoStreamer) handleRangeRequest(w http.ResponseWriter, r *http.Request, file *os.File, fileSize int64, rangeHeader string) error {
	// Parse range header (e.g., "bytes=0-1023")
	rangeStr := strings.TrimPrefix(rangeHeader, "bytes=")
	parts := strings.Split(rangeStr, "-")

	if len(parts) != 2 {
		return fmt.Errorf("invalid range header")
	}

	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid start range")
	}

	var end int64
	if parts[1] == "" {
		end = fileSize - 1
	} else {
		end, err = strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid end range")
		}
	}

	// Validate range
	if start < 0 || end >= fileSize || start > end {
		w.Header().Set("Content-Range", fmt.Sprintf("bytes */%d", fileSize))
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
		return fmt.Errorf("invalid range")
	}

	// Set range response headers
	contentLength := end - start + 1
	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
	w.Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))
	w.WriteHeader(http.StatusPartialContent)

	// Seek to start position
	_, err = file.Seek(start, 0)
	if err != nil {
		return fmt.Errorf("error seeking to position")
	}

	// Stream the requested range
	_, err = io.CopyN(w, file, contentLength)
	if err != nil {
		return fmt.Errorf("error streaming range")
	}

	return nil
}

// GetVideoInfo returns information about a video file
func (vs *VideoStreamer) GetVideoInfo(videoID string) (*VideoInfo, error) {
	videoPath := filepath.Join(vs.videoDirectory, videoID+".mp4")

	fileInfo, err := os.Stat(videoPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("video not found")
		}
		return nil, fmt.Errorf("error accessing video file")
	}

	return &VideoInfo{
		ID:       videoID,
		Size:     fileInfo.Size(),
		Modified: fileInfo.ModTime().Format("2006-01-02 15:04:05"),
		Path:     videoPath,
	}, nil
}

// VideoInfo represents information about a video file
type VideoInfo struct {
	ID       string `json:"id"`
	Size     int64  `json:"size"`
	Modified string `json:"modified"`
	Path     string `json:"path"`
}
