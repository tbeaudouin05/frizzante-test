import type { HistoryEntry, View } from "$lib/scripts/core/types.ts"

let lastUrl: false | string = false

export async function swap(
    target: HTMLAnchorElement | HTMLFormElement,
    view: View<unknown>,
): Promise<() => void> {
    if (lastUrl === false) {
        lastUrl = location.toString()
    }

    let res: Response
    let method: "GET" | "POST" = "GET"
    const body: Record<string, string> = {}

    if (target.nodeName === "A") {
        const anchor = target as HTMLAnchorElement
        res = await fetch(anchor.href, {
            headers: {
                Accept: "application/json",
            },
            credentials: "same-origin",
        })
    } else if (target.nodeName === "FORM") {
        const form = target as HTMLFormElement
        const data = new FormData(form)
        const params = new URLSearchParams()
        let query = ""

        data.forEach(function each(value, key) {
            if (value instanceof File) {
                return
            }
            body[key] = `${value}`
            params.append(key, `${value}`)
        })

        const declared = (form.getAttribute("method") || "GET").toUpperCase()
        method = (declared === "POST" ? "POST" : "GET")

        if (method === "GET") {
            query = `${params.toString()}`
            if (query !== "") {
                if (form.action.includes("?")) {
                    query = "&" + query
                } else {
                    query = "?" + query
                }
            }
            res = await fetch(`${form.action}${query}`, {
                headers: {
                    Accept: "application/json",
                },
                credentials: "same-origin",
            })
        } else {
            res = await fetch(form.action, {
                method,
                body: data,
                headers: {
                    Accept: "application/json",
                },
                credentials: "same-origin",
            })
        }
    } else {
        return function push() {}
    }

    // Follow server-side navigate (302 with Location header) by refetching JSON
    if (res.status >= 300 && res.status < 400) {
        const loc = res.headers.get("Location")
        if (loc) {
            res = await fetch(loc, {
                headers: { Accept: "application/json" },
                credentials: "same-origin",
            })
        }
    }

    const txt = await res.text()

    if ("" === txt) {
        return function push() {}
    }

    const remote = JSON.parse(txt)

    view.align = remote.align
    view.name = remote.name
    view.render = remote.render
    if (view.align === 1) {
        if (typeof view.props != "object") {
            console.warn(
                "view alignment intends to merge props, but local view props is not an object",
            )
            // Noop.
        } else if (typeof remote.props != "object") {
            console.warn(
                "view alignment intends to merge props, but remote props is not an object",
            )
            // Noop.
        } else {
            view.props = {
                ...view.props,
                ...remote.props,
            }
        }
    } else {
        view.props = remote.props
    }

    const stationary = lastUrl === res.url
    lastUrl = res.url

    return function push() {
        if (stationary) {
            return
        }

        const entry: HistoryEntry = {
            nodeName: target.nodeName,
            method,
            url: res.url,
            body,
        }

        window.history.pushState(JSON.stringify(entry), "", res.url)
    }
}
