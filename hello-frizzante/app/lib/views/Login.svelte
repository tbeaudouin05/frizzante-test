<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.ts";
  import { Input } from "$lib/components/ui/input/index.ts";
  import { supabase } from "$lib/supabaseClient.ts";

  let email = "";
  let password = "";
  let loading = false;
  let errorMsg = "";

  async function signIn(e: Event) {
    e.preventDefault();
    errorMsg = "";
    loading = true;
    try {
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
</script>

<div class="mx-auto max-w-md p-6">
  <h1 class="text-2xl font-semibold mb-4">Login</h1>
  {#if errorMsg}
    <div class="mb-3 text-red-500 text-sm">{errorMsg}</div>
  {/if}
  <form on:submit={signIn} class="space-y-3">
    <div>
      <label class="block text-sm mb-1" for="email">Email</label>
      <Input id="email" type="email" required bind:value={email} placeholder="you@example.com" />
    </div>
    <div>
      <label class="block text-sm mb-1" for="password">Password</label>
      <Input id="password" type="password" required bind:value={password} />
    </div>
    <Button type="submit" disabled={loading} class="w-full">{loading ? "Signing in..." : "Sign in"}</Button>
  </form>
</div>
