<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.ts";
  import { href } from "$lib/scripts/core/href.ts";
  import ThemeToggle from "$lib/components/ThemeToggle.svelte";
</script>

<nav class="w-full border-b border-border/80 bg-background/80 backdrop-blur supports-[backdrop-filter]:bg-background/60">
  <div class="mx-auto max-w-6xl px-6 h-14 flex items-center justify-between gap-3">
    <a class="font-semibold tracking-tight text-foreground text-lg bg-gradient-to-r from-primary to-foreground/70 bg-clip-text text-transparent" {...href("/")}>Frizzante</a>
    <div class="flex items-center gap-2">
      <Button variant="ghost" class="rounded-full" {...href("/")}>Home</Button>
      <Button variant="ghost" class="rounded-full" {...href("/todos")}>Todos</Button>
      <Button variant="ghost" class="rounded-full" {...href("/lessons")}>Lessons</Button>
      <ThemeToggle />
    </div>
  </div>
</nav>
