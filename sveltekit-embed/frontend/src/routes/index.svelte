<script>
  import { onMount } from "svelte";
  import { getIP } from "$lib/ip";
  import Title from "$lib/Title.svelte";
  let detail;
  onMount(async () => {
    detail = await getIP();
  });
</script>

<svelte:head>
  <title>Fiber SvelteKit</title>
</svelte:head>

<Title title="Your IP Address" />
<div class="py-8 px-4 font-bold uppercase text-4xl font-mono">
  {detail ? detail.ip : "-"}
</div>
