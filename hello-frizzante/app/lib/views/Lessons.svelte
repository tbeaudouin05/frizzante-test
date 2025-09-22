<script lang="ts">
  import Layout from "$lib/components/Layout.svelte"
  import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from "$lib/components/ui/card/index.ts"
  import { Input } from "$lib/components/ui/input/index.ts"
  import { Button } from "$lib/components/ui/button/index.ts"
  import { Separator } from "$lib/components/ui/separator/index.ts"
  import { Calendar } from "$lib/components/ui/calendar/index.ts"
  import { Select, SelectTrigger, SelectContent, SelectGroup, SelectLabel, SelectItem } from "$lib/components/ui/select/index.ts"
  import { getLocalTimeZone, today, type CalendarDate } from "@internationalized/date"
  import { action } from "$lib/scripts/core/action.ts"

  // Props from server
  let { lessons = [], error }: any = $props()

  let date: CalendarDate | undefined = $state(today(getLocalTimeZone()))
  let time: string = $state("10:00") // default time slot
  let student: string = $state("")

  const dateISO = $derived.by(() =>
    date ? `${date.year}-${String(date.month).padStart(2, "0")}-${String(date.day).padStart(2, "0")}` : ""
  )

  const timeSlots = Array.from({ length: (18 - 7 + 1) }, (_, i) => {
    const h = 7 + i // 07:00 through 18:00
    return `${String(h).padStart(2, "0")}:00`
  })

  function onBook() {
    const selected = date
    const d = selected ? selected.toDate(getLocalTimeZone()) : new Date()
    alert(`Booked lesson on ${d.toDateString()} at ${time}`)
  }
</script>

<Layout title="Book a Surf Lesson">
  <main class="min-h-screen">
    <section class="container mx-auto max-w-3xl px-6 py-10">
      <div class="mb-8 text-center">
        <h1 class="text-4xl font-bold tracking-tight">Book a Surf Lesson</h1>
        <p class="mt-2 text-muted-foreground">Pick your preferred date and time</p>
      </div>

      <Card class="border-border/60">
        <CardHeader>
          <CardTitle>New Booking</CardTitle>
          <CardDescription>Fill in your name, select a date and time, then book.</CardDescription>
        </CardHeader>
        <CardContent class="space-y-4">
          <form {...action("/lessons/book")} method="GET" class="grid gap-4 sm:grid-cols-2">
            <div class="sm:col-span-2">
              <label for="studentName" class="mb-1 block text-sm font-medium">Student name</label>
              <Input id="studentName" name="student" bind:value={student} placeholder="e.g. Kelly Slater" required />
            </div>

            <div>
              <fieldset class="rounded-md border p-2">
                <legend class="px-1 text-sm font-medium">Date</legend>
                <Calendar type="single" bind:value={date} />
              </fieldset>
              <input type="hidden" name="date" value={dateISO} />
            </div>

            <div>
              <fieldset class="rounded-md">
                <legend id="timeLabel" class="mb-1 block text-sm font-medium">Time</legend>
                <div role="group" aria-labelledby="timeLabel">
                  <Select type="single" bind:value={time}>
                    <SelectTrigger class="w-full" aria-labelledby="timeLabel">
                      <span data-slot="select-value">{time || 'Select a time'}</span>
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectLabel>Available slots</SelectLabel>
                        {#each timeSlots as t}
                          <SelectItem value={t}>{t}</SelectItem>
                        {/each}
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                </div>
              </fieldset>
              <input type="hidden" name="time" value={time} />
            </div>

            <div class="sm:col-span-2 flex items-center justify-between">
              <div class="text-sm text-muted-foreground">
                {#if dateISO && time}
                  Booking for <span class="font-medium">{dateISO}</span> at <span class="font-medium">{time}</span>
                {:else}
                  Select a date and a time
                {/if}
              </div>
              <Button type="submit" disabled={!dateISO || !time || !student.trim()}>Book</Button>
            </div>
          </form>
        </CardContent>
        <CardFooter>
          <p class="text-sm text-muted-foreground">We’ll confirm your booking instantly.</p>
        </CardFooter>
      </Card>

      <div class="mt-10">
        <h2 class="mb-3 text-xl font-semibold">Upcoming Lessons</h2>
        <Separator class="mb-4" />
        {#if lessons.length === 0}
          <p class="text-muted-foreground">No bookings yet.</p>
        {:else}
          <div class="space-y-2">
            {#each lessons as l, i (i)}
              <div class="flex items-center justify-between rounded-lg border bg-card p-3 text-card-foreground">
                <div>
                  <div class="font-medium">{l.student}</div>
                  <div class="text-sm text-muted-foreground">{l.date} at {l.time}</div>
                </div>
                <form {...action("/lessons/cancel")} method="GET">
                  <input type="hidden" name="index" value={i} />
                  <Button type="submit" variant="ghost">Cancel</Button>
                </form>
              </div>
            {/each}
          </div>
        {/if}
      </div>
    </section>
  </main>
</Layout>
