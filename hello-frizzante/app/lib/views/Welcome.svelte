<script lang="ts">
    import Layout from "$lib/components/Layout.svelte"
    import { href } from "$lib/scripts/core/href.ts"
    import { fade } from "svelte/transition"
    import { Button } from "$lib/components/ui/button/index.ts"
    import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "$lib/components/ui/card/index.ts"
    import { Rocket, Sparkles, Gauge } from "@lucide/svelte"
</script>

<Layout title="Welcome">
    <main
        in:fade={{ duration: 300 }}
        class="min-h-screen flex items-center justify-center p-4"
    >
        <div class="max-w-6xl w-full space-y-12">
            <section class="relative overflow-hidden rounded-2xl border bg-card/80 backdrop-blur shadow-xl p-8 md:p-12">
                <div class="absolute inset-0 pointer-events-none opacity-50"
                    style="background:
                        radial-gradient(600px 200px at 10% 0%, hsl(var(--primary)/.15), transparent),
                        radial-gradient(600px 200px at 90% 20%, hsl(var(--ring)/.15), transparent);"></div>
                <div class="relative text-center space-y-4">
                    <h1 class="text-5xl md:text-7xl font-semibold tracking-tight bg-gradient-to-r from-primary to-foreground/70 bg-clip-text text-transparent">
                        Frizzante
                    </h1>
                    <p class="text-lg md:text-xl text-muted-foreground max-w-3xl mx-auto">
                        Modern Go + Svelte framework with server-rendered views and a sleek shadcn UI foundation.
                    </p>
                    <div class="flex items-center justify-center gap-3 pt-1">
                        <Button size="lg" {...href("/todos")}>Get Started</Button>
                        <Button variant="outline" size="lg" href="https://razshare.github.io/frizzante-docs/guides/get-started" target="_blank">Docs</Button>
                    </div>
                </div>
            </section>

            <section class="grid md:grid-cols-3 gap-4">
                <Card class="shadow-lg border-2">
                    <CardHeader class="space-y-2">
                        <div class="inline-flex size-10 items-center justify-center rounded-full bg-primary/10 text-primary">
                            <Gauge size={18} />
                        </div>
                        <CardTitle>Lightning Fast</CardTitle>
                        <CardDescription>Built with Go for exceptional performance</CardDescription>
                    </CardHeader>
                    <CardContent class="text-sm text-muted-foreground">Hot reloading, SSR by default, and optimized assets.
                    </CardContent>
                </Card>
                <Card class="shadow-lg border-2">
                    <CardHeader class="space-y-2">
                        <div class="inline-flex size-10 items-center justify-center rounded-full bg-primary/10 text-primary">
                            <Sparkles size={18} />
                        </div>
                        <CardTitle>Reactive UI</CardTitle>
                        <CardDescription>Powered by Svelte for smooth interfaces</CardDescription>
                    </CardHeader>
                    <CardContent class="text-sm text-muted-foreground">Use modern Svelte 5 runes and shadcn components.
                    </CardContent>
                </Card>
                <Card class="shadow-lg border-2">
                    <CardHeader class="space-y-2">
                        <div class="inline-flex size-10 items-center justify-center rounded-full bg-primary/10 text-primary">
                            <Rocket size={18} />
                        </div>
                        <CardTitle>Developer First</CardTitle>
                        <CardDescription>Opinionated choices that streamline workflow</CardDescription>
                    </CardHeader>
                    <CardContent class="text-sm text-muted-foreground">Simple routing, typed props, and ergonomic DX.
                    </CardContent>
                </Card>
            </section>
        </div>
    </main>
</Layout>
