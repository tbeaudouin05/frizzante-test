<style>
    .app-bg {
        position: fixed;
        inset: 0;
        pointer-events: none;
        background:
            radial-gradient(800px 320px at 50% -10%, hsl(var(--primary) / 0.10), transparent),
            radial-gradient(700px 260px at 100% 0%, hsl(var(--secondary) / 0.08), transparent),
            linear-gradient(to bottom right, hsl(var(--background)), hsl(var(--background)));
        z-index: -1;
    }
    .app-font {
        font-family: Inter, ui-sans-serif, system-ui, -apple-system, Segoe UI,
            Roboto, Helvetica, Arial, "Apple Color Emoji", "Segoe UI Emoji";
    }
</style>

<script lang="ts">
    import { getContext, type Snippet } from "svelte"
    import type { View } from "$lib/scripts/core/types.ts"
    import Navbar from "./Navbar.svelte"

    type Props = {
        title?: string
        children: Snippet
    }

    const view = getContext("view") as View<unknown>

    let { title = view.name, children }: Props = $props()
</script>

<svelte:head>
    <meta charset="UTF-8" />
    <meta
        name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0"
    />
    <title>{title}</title>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link
        href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap"
        rel="stylesheet"
    />
</svelte:head>

<div class="app-font text-foreground">
    <div class="app-bg"></div>
    <header>
        <Navbar />
    </header>
    <div class="mx-auto max-w-6xl px-6 py-14 md:py-20">
        {@render children()}
    </div>
    <footer class="py-8 text-center text-sm text-muted-foreground">
        Built with Frizzante, Svelte 5 and shadcn-style components
    </footer>
    
</div>
