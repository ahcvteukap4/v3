package cli

import "os"

// ... existing code ...

func (a *App) Run(arguments []string) error {
    // ... existing parsing logic ...

    if a.Before != nil {
        if err := a.Before(ctx); err != nil {
            return err
        }
        // Re-evaluate env vars for flags not set by CLI
        for _, f := range a.Flags {
            if !ctx.IsSet(f.Names()[0]) {
                if envVar, ok := f.(interface{ ApplyEnvVars() error }); ok {
                    _ = envVar.ApplyEnvVars()
                }
            }
        }
    }

    return a.Action(ctx)
}