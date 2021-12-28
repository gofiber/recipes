<script>
  import "../tailwind.css";
  import Theme from "$lib/Theme.svelte";
  import Nav from "$lib/Nav.svelte";
</script>

<div class="absolute top-4 right-4">
  <Theme />
</div>
<main class="flex items-center justify-center min-h-screen p-4">
  <div
    class="w-full max-w-md text-center rounded bg-white dark:bg-slate-800 shadow-2xl shadow-slate-400 dark:shadow-black"
  >
    <slot />
    <Nav />
  </div>
</main>
