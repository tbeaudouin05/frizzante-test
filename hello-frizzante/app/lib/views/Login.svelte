<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.ts";
  import { Input } from "$lib/components/ui/input/index.ts";
  import { onMount } from "svelte";

  let email = "";
  let password = "";
  let loading = false;
  let errorMsg = "";
  let supabase: any = null;
  let clientReady = false;
  let isSignup = false;

  onMount(async () => {
    const mod = await import("$lib/supabaseClient.ts");
    supabase = mod.supabase;
    clientReady = true;
  });

  async function signIn(e: Event) {
    e.preventDefault();
    errorMsg = "";
    loading = true;
    try {
      if (!clientReady || !supabase) {
        // Client not initialized yet — ignore submit.
        return;
      }
      const { data, error } = await supabase.auth.signInWithPassword({ email, password });
      if (error) {
        throw error;
      }
      const access = data.session?.access_token;
      if (!access) {
        throw new Error("No access token returned by Supabase");
      }
      // Handshake with backend to set HttpOnly cookie
      const resp = await fetch("/auth/session", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${access}`,
          Accept: "application/json",
        },
        credentials: "same-origin",
      });
      if (!resp.ok && resp.status !== 204) {
        const txt = await resp.text();
        throw new Error(txt || "Failed to establish session");
      }
      // Navigate home
      window.location.href = "/";
    } catch (err: any) {
      errorMsg = err?.message || String(err);
    } finally {
      loading = false;
    }
  }

  async function signUp(e: Event) {
    e.preventDefault();
    errorMsg = "";
    loading = true;
    try {
      if (!clientReady || !supabase) return;
      const { data, error } = await supabase.auth.signUp({ email, password });
      if (error) throw error;
      const access = data.session?.access_token;
      if (!access) {
        // Most projects require email confirmation: no session until confirmed
        errorMsg = "Check your email to confirm your account, then come back to sign in.";
        return;
      }
      const resp = await fetch("/auth/session", {
        method: "POST",
        headers: { Authorization: `Bearer ${access}`, Accept: "application/json" },
        credentials: "same-origin",
      });
      if (!resp.ok && resp.status !== 204) {
        const txt = await resp.text();
        throw new Error(txt || "Failed to establish session");
      }
      window.location.href = "/";
    } catch (err: any) {
      errorMsg = err?.message || String(err);
    } finally {
      loading = false;
    }
  }
</script>

<div class="mx-auto max-w-md p-6">
  <h1 class="text-2xl font-semibold mb-4">{isSignup ? 'Create account' : 'Login'}</h1>
  {#if !clientReady}
    <div class="mb-3 text-amber-600 text-sm">Initializing auth…</div>
  {/if}
  {#if errorMsg}
    <div class="mb-3 text-red-500 text-sm">{errorMsg}</div>
  {/if}
  <form on:submit={isSignup ? signUp : signIn} class="space-y-3">
    <div>
      <label class="block text-sm mb-1" for="email">Email</label>
      <Input id="email" type="email" required bind:value={email} placeholder="you@example.com" />
    </div>
    <div>
      <label class="block text-sm mb-1" for="password">Password</label>
      <Input id="password" type="password" required bind:value={password} />
    </div>
    <Button type="submit" disabled={loading || !clientReady} class="w-full">
      {#if loading}
        {isSignup ? 'Creating account...' : 'Signing in...'}
      {:else}
        {isSignup ? 'Create account' : 'Sign in'}
      {/if}
    </Button>
    <div class="text-xs text-muted-foreground mt-2">
      {#if isSignup}
        Already have an account?
        <a href="#" on:click|preventDefault={() => { isSignup = false; errorMsg = ''; }} class="underline">Sign in</a>
      {:else}
        New here?
        <a href="#" on:click|preventDefault={() => { isSignup = true; errorMsg = ''; }} class="underline">Create an account</a>
      {/if}
    </div>
  </form>
</div>
