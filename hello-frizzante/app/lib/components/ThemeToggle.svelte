<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.ts";
  let isDark = $state(false);
  function apply(mode: boolean) {
    const root = document.documentElement;
    if (mode) root.classList.add("dark");
    else root.classList.remove("dark");
    isDark = mode;
    try { localStorage.setItem("theme:dark", mode ? "1" : "0"); } catch {}
  }

  function toggle() {
    apply(!isDark);
  }

  $effect.pre(() => {
    // Initialize based on prefers-color-scheme
    let startDark = false;
    try { startDark = localStorage.getItem("theme:dark") === "1"; } catch {}
    if (startDark === false) {
      const prefersDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
      startDark = prefersDark;
    }
    apply(startDark);
  });
</script>

<Button variant="ghost" onclick={toggle} aria-label="Toggle theme">
  {#if isDark}
    Light
  {:else}
    Dark
  {/if}
</Button>
