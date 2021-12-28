<script context="module">
  export function load({ error, status }) {
    return {
      props: { error, status },
    };
  }
</script>

<script>
  import { dev } from "$app/env";

  export let error, status;
</script>

<svelte:head>
  <title>{status}</title>
</svelte:head>

<div class="p-4">
  <h1 class="text-8xl font-bold uppercase mb-4">{status}</h1>

  <p class="mb-4">{error.message}</p>

  <p>
    <a
      href="/"
      class="py-2 px-4 inline-block rounded text-white bg-primary hover:bg-opacity-75 focus:bg-opacity-75"
    >
      Return Home
    </a>
  </p>

  {#if dev && error.stack}
    <pre
      class="w-full overflow-x-auto p-2 bg-black bg-opacity-5 dark:bg-opacity-25 mt-4 rounded">{error.stack}</pre>
  {/if}
</div>
