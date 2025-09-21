import { defineConfig } from "vite"
import { svelte } from "@sveltejs/vite-plugin-svelte"
import tailwindcss from "@tailwindcss/vite"

const IS_DEV = (process.env.DEV ?? "0") === "1"

let sourcemap: "inline" | boolean = false
if (IS_DEV) {
    sourcemap = "inline"
}

// https://vite.dev/config/
export default defineConfig({
    base: "./",
    plugins: [
        tailwindcss(),
        svelte({
            compilerOptions: {
                css: "injected",
            },
        }),
    ],
    resolve: {
        alias: {
            $lib: "./lib",
        },
    },
    build: {
        copyPublicDir: false,
        sourcemap,
        rollupOptions: {
            input: {
                index: "./index.html",
            },
        },
    },
})
