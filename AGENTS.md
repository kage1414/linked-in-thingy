# Agent Instructions for Job Board Project

## Project Overview

This is a job board application with a Go backend (REST API) and React TypeScript frontend, featuring video streaming capabilities and PostgreSQL database integration.

## Architecture Patterns

### Component Structure

- **Container/Visual Pattern**: All React components follow a container/visual separation
  - `ComponentName.container.tsx` - Business logic, state management, data processing
  - `ComponentName.tsx` - Visual presentation, UI rendering
  - `index.ts` - Clean exports
  - `ComponentName.css` - Component-specific styling (when needed)

### File Organization

- Keep related files together in component directories
- Separate CSS files from component logic
- Use clean, descriptive file names
- Frontend code belongs in ./frontend
- Backend code belongs in ./backend
- Database code belonds in ./backend/database
- main.go, go.sum, and go.mod are the only backend-related files allowed in the top level directory. everything else belongs in ./backend
- Maintain consistent directory structure
- Global reusable components are contained in ./frontend/src/components
- Do not move into new directory unless asked. If rejected, component should exist in frontend/src/components
- Business logic should be contained in [ComponentName].container.tsx
- Visual logic should be contained in [ComponentName].tsx
- Locally used functions should remain in their respective component folder
- Globally used functions should exist in frontend/src
- Local custom hooks should remain in hooks.tsx in their respective component folder.
- Global custom hooks should exist in frontend/src/hooks
- Global components should exist at the top level in frontend/src/components
- No folders with a single folder as their contents
- Global components are prefixed with LIT[ComponentName].tsx
- Global components exist in frontend/src/components/LITComponents
- "hooks", "assets", and "features" are the only directories allowed under frontend/src

## Development Guidelines

### Code Quality

- **TypeScript**: Use strict typing, define interfaces for all props and data structures
- **Error Handling**: Implement proper error boundaries and user-friendly error messages
- **Loading States**: Always provide loading indicators for async operations
- **Accessibility**: Include proper ARIA attributes and semantic HTML
- Test files should be named [ComponentName].test.tsx for React files.
- Test files should be named [ComponentName].test.ts for non-react files.
- Helper functions should be in utils.ts. If only used locally, they should be in their component folder. If used globally, they should be in frontend/src/utils.ts
- If I specify a specific file, only update that file. You may fix imports if necessary. Please ask before modifying another file
- Don't create empty files
- Don't use useMemo
- Don't use deprecated methods
- Components with no business logic should use an implicit return and omit the return statement

### Dependencies

- Use libraries first before building helper functions.
- Ask before importing.
- Prefer packages with <5 dependencies
- Prefer popular and vetted packages, > 1000000 weekly installations

### Component Design

- **Single Responsibility**: Each component should have one clear purpose
- **Reusability**: Design components to be reusable across the application
- **Props Interface**: Always define clear TypeScript interfaces for component props
- **Default Props**: Use sensible defaults for optional props

### State Management

- Keep state as close to where it's used as possible
- Use proper React hooks (useState, useEffect, useCallback, useMemo)
- Handle loading, error, and success states consistently
- Implement proper cleanup for side effects
- Any function using 3 or more hooks should be refactored into a custom hook
- Custom hooks should have their own file

## Backend Guidelines

### API Design

- Use RESTful conventions for endpoints
- Implement proper HTTP status codes
- Include comprehensive error handling
- Use consistent response formats

### Database

- Use GORM for database operations
- Implement proper migrations
- Include data validation
- Handle database errors gracefully

## Testing & Quality Assurance

### Before Committing

- Run `npm run build` to ensure no compilation errors
- Check for linting errors with `read_lints`
- Verify all imports are correct
- Test that the application builds successfully

### Code Review Checklist

- [ ] TypeScript types are properly defined
- [ ] Error handling is implemented
- [ ] Loading states are handled
- [ ] Components follow container/visual pattern
- [ ] CSS is properly organized
- [ ] No console errors or warnings

## File Management

### When Creating New Components

1. Create component directory: `components/ComponentName/`
2. Add container file: `ComponentName.container.tsx`
3. Add visual file: `ComponentName.tsx`
4. Add index file: `index.ts`
5. Add CSS file if needed: `ComponentName.css`

### When Refactoring

- Maintain existing functionality
- Update all imports when moving files
- Test thoroughly after changes
- Keep git history clean with descriptive commits

## Communication Style

### When Working on Tasks

- Break down complex tasks into smaller, manageable steps
- Use TODO lists to track progress
- Provide clear explanations of changes made
- Ask for clarification when requirements are unclear

### When Reporting Issues

- Provide specific error messages
- Include relevant file paths and line numbers
- Suggest potential solutions
- Test fixes before reporting success

## Project-Specific Rules

### Frontend

- Use functional components with hooks
- Implement proper error boundaries
- Use the Button component for all interactive elements
- Follow the established CSS class naming conventions

### Backend

- Use Gin framework for HTTP routing
- Implement proper CORS configuration
- Use GORM for database operations
- Include proper logging for debugging

### Database

- Use PostgreSQL with Docker Compose
- Implement proper migrations
- Include sample data seeding
- Handle connection errors gracefully

## Common Commands

### Development

```bash
# Start development servers
npm run dev

# Build frontend
npm run frontend:build

# Start database
make db-up

# Stop database
make db-down
```

### Testing

```bash
# Build and test
cd frontend && npm run build

# Check for linting errors
# (Use read_lints tool)
```

## Error Handling Patterns

### Frontend Errors

- Display user-friendly error messages
- Provide retry mechanisms where appropriate
- Log errors to console for debugging
- Implement proper loading states

### Backend Errors

- Return appropriate HTTP status codes
- Include descriptive error messages
- Log errors for debugging
- Handle database connection issues

## Performance Considerations

### Frontend

- Use React.memo for expensive components
- Implement proper key props for lists
- Optimize bundle size
- Use lazy loading where appropriate

### Backend

- Implement proper database indexing
- Use connection pooling
- Optimize database queries
- Handle concurrent requests properly

## Security Guidelines

### API Security

- Implement proper CORS configuration
- Validate all input data
- Use parameterized queries
- Implement rate limiting where needed

### Frontend Security

- Sanitize user input
- Use HTTPS in production
- Implement proper authentication flows
- Handle sensitive data appropriately

---

_This file should be updated as the project evolves and new patterns emerge._
