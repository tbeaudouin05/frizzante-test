<script lang="ts">
  import Layout from "$lib/components/Layout.svelte"
  import Navbar from "$lib/components/Navbar.svelte"
  import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "$lib/components/ui/card/index.ts"
  import { Input } from "$lib/components/ui/input/index.ts"
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

  const timeSlots = Array.from({ length: (20 - 6) * 2 + 1 }, (_, i) => {
    const h = Math.floor(6 + i / 2)
    const m = i % 2 === 0 ? "00" : "30"
    return `${String(h).padStart(2, "0")}:${m}`
  })

  function onBook() {
    const selected = date
    const d = selected ? selected.toDate(getLocalTimeZone()) : new Date()
    alert(`Booked lesson on ${d.toDateString()} at ${time}`)
  }
</script>

<Layout title="Book a Surf Lesson">
  <main class="min-h-screen">
    <Navbar />

    <section class="mx-auto max-w-3xl px-6 py-10 space-y-8">
      <div class="text-center space-y-2">
        <h1 class="text-3xl md:text-4xl font-semibold tracking-tight">Book a Surf Lesson</h1>
        <p class="text-muted-foreground">Choose a date and time for your session</p>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>New Booking</CardTitle>
          <CardDescription>Fill in your name, select a date and time, then book.</CardDescription>
        </CardHeader>
        <CardContent class="space-y-6">
          <form {...action("/lessons/book")} method="GET" class="space-y-6">
            <!-- Student name -->
            <div class="space-y-2">
              <label for="student-name" class="text-sm text-muted-foreground">Student name</label>
              <Input id="student-name" name="student" bind:value={student} placeholder="e.g. Kelly Slater" class="w-full" />
            </div>

            <!-- Date + Time grid -->
            <div class="grid md:grid-cols-2 gap-6 items-start">
              <div class="space-y-2">
                <label id="date-label" class="text-sm text-muted-foreground">Date</label>
                <div class="flex justify-center md:justify-start" aria-labelledby="date-label">
                  <Calendar type="single" bind:value={date} class="rounded-lg border shadow-sm" />
                </div>
              </div>
              <div class="space-y-2">
                <label for="lesson-time" class="text-sm text-muted-foreground">Time</label>
                <!-- shadcn Select time picker -->
                <Select type="single" bind:value={time}>
                  <SelectTrigger id="lesson-time" class="w-full md:w-[220px]">
                    {time || "Select a time"}
                  </SelectTrigger>
                  <SelectContent class="w-[220px]">
                    <SelectGroup>
                      <SelectLabel>Available slots</SelectLabel>
                      {#each timeSlots as t}
                        <SelectItem value={t}>{t}</SelectItem>
                      {/each}
                    </SelectGroup>
                  </SelectContent>
                </Select>
                <!-- keep time value submitted in form -->
                <input type="hidden" name="time" value={time} />
              </div>
            </div>

            <input type="hidden" name="date" value={dateISO} />

            <div class="text-sm text-muted-foreground">Select a date and a time</div>

            <div class="flex justify-end pt-2">
              <button type="submit" class="h-10 px-4 rounded-md bg-primary text-primary-foreground hover:bg-primary/90">
                Book
              </button>
            </div>

            <div class="text-sm text-muted-foreground">We’ll confirm your booking instantly.</div>
          </form>
        </CardContent>
      </Card>

      <section class="space-y-3">
        <h2 class="font-semibold">Upcoming Lessons</h2>
        {#if lessons.length === 0}
          <div class="text-sm text-muted-foreground">No bookings yet.</div>
        {:else}
          <div class="space-y-2">
            {#each lessons as l, i (i)}
              <div class="flex items-center justify-between rounded-md border bg-card p-3">
                <div>
                  <div class="font-medium">{l.student}</div>
                  <div class="text-sm text-muted-foreground">{l.date} at {l.time}</div>
                </div>
                <form {...action("/lessons/cancel")} method="GET">
                  <input type="hidden" name="index" value={i} />
                  <button type="submit" class="h-9 px-3 rounded-md border bg-background hover:bg-accent">Cancel</button>
                </form>
              </div>
            {/each}
          </div>
        {/if}
      </section>
    </section>
  </main>
</Layout>
