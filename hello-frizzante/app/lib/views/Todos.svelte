<script lang="ts">
    import Icon from "$lib/components/icons/Icon.svelte"
    import Layout from "$lib/components/Layout.svelte"
    import { action } from "$lib/scripts/core/action.ts"
    import { href } from "$lib/scripts/core/href.ts"
    import { mdiArrowLeft, mdiCheck, mdiDelete, mdiPlus } from "@mdi/js"
    import { fade, slide } from "svelte/transition"
    import type {Props} from "$lib/types/gen/main/lib/routes/handlers/todos/Props";
    import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "$lib/components/ui/card/index.ts"
    import { Input } from "$lib/components/ui/input/index.ts"
    import { Button } from "$lib/components/ui/button/index.ts"
    import { Alert } from "$lib/components/ui/alert/index.ts"

    let { todos = [], error }: Props = $props()
</script>

<Layout title="Todos">
    <main
        in:fade={{ duration: 300 }}
        class="min-h-screen flex items-center justify-center p-4"
    >
        <div class="w-full min-w-[360px] max-w-2xl space-y-8">
            <div class="relative text-center space-y-2 mb-1">
                <Button variant="ghost" class="absolute left-0 -top-1 text-sm gap-1" {...href("/")}>
                    <Icon path={mdiArrowLeft} size="16" /> Back
                </Button>
                <h1 class="text-3xl md:text-4xl font-semibold tracking-tight">My Tasks</h1>
                <p class="text-base text-muted-foreground">Organize and track your daily activities</p>
            </div>

            <Card>
                <CardHeader class="pb-3">
                    <CardTitle class="text-xl">Add Task</CardTitle>
                    <CardDescription>Quickly add a new task to your list</CardDescription>
                </CardHeader>
                <CardContent class="space-y-4">
                    <form {...action("/add")} method="GET" class="flex gap-2">
                        <Input name="description" placeholder="Add a new task..." required class="flex-1" />
                        <Button type="submit" class="min-w-24">
                            <Icon path={mdiPlus} size="18" /> Add
                        </Button>
                    </form>

                    {#if error}
                        <div in:slide>
                            <Alert variant="destructive">{error}</Alert>
                        </div>
                    {/if}
                </CardContent>
            </Card>

            <Card>
                <CardHeader class="pb-2">
                    <CardTitle class="text-xl">Your Tasks</CardTitle>
                    <CardDescription>Click the checkbox to mark a task as done</CardDescription>
                </CardHeader>
                <CardContent class="space-y-2">
                    {#if todos.length === 0}
                        <div class="text-center py-6 text-muted-foreground">No tasks yet. Add one above to get started!</div>
                    {:else}
                        <div class="space-y-2">
                            {#each todos as todo, index (index)}
                                <div in:slide class="group flex items-center gap-3 p-3 rounded-lg border hover:border-primary/50 hover:shadow-sm transition-all bg-card">
                                    <form
                                        {...action(todo.checked ? "/uncheck" : "/check")}
                                        method="GET"
                                        class="flex-shrink-0"
                                    >
                                        <input type="hidden" name="index" value={index} />
                                        <button
                                            type="submit"
                                            class="flex h-4 w-4 items-center justify-center rounded-sm border border-primary text-primary shadow focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground"
                                            data-state={todo.checked ? "checked" : "unchecked"}
                                            role="checkbox"
                                            aria-checked={todo.checked}
                                            aria-label={todo.checked ? "Uncheck" : "Check"}
                                            on:click={(e) => { e.preventDefault(); todo.checked = !todo.checked; (e.currentTarget as HTMLButtonElement).form?.requestSubmit(); }}
                                        >
                                            {#if todo.checked}
                                                <Icon path={mdiCheck} size="12" />
                                            {/if}
                                        </button>
                                    </form>

                                    <span class={`flex-1 ${todo.checked ? "line-through text-muted-foreground" : "text-foreground"}`}>
                                        {todo.description}
                                    </span>

                                    <form {...action("/remove")} method="GET" class="flex-shrink-0">
                                        <input type="hidden" name="index" value={index} />
                                        <Button type="submit" variant="ghost" aria-label="Delete" class="text-destructive">
                                            <Icon path={mdiDelete} size="18" />
                                        </Button>
                                    </form>
                                </div>
                            {/each}
                        </div>
                        <div class="pt-2 text-sm text-muted-foreground text-center">
                            {todos.filter(t => !t.checked).length} of {todos.length} tasks remaining
                        </div>
                    {/if}
                </CardContent>
            </Card>
        </div>
    </main>
</Layout>
