<script>
  import { onMount } from "svelte";
  import { getByIP } from "$lib/ip";
  import Title from "$lib/Title.svelte";
  import Detail from "$lib/Detail.svelte";
  let ip,
    detail,
    error,
    input,
    loading = false,
    timer = 0;

  async function getIPDetail() {
    clearTimeout(timer);
    error = null;
    detail = null;
    if (!ip) {
      return;
    }
    loading = true;
    const ipformat =
      /^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/;
    if (ip.match(ipformat)) {
      detail = await getByIP(ip);
    } else {
      error = "invalid ip format";
    }
    loading = false;
    timer = setTimeout(() => {
      error = null;
    }, 5000);
  }

  onMount(() => {
    setTimeout(() => {
      input.focus();
    }, 250);
  });
</script>

<svelte:head>
  <title>Search IP - Fiber SvelteKit</title>
</svelte:head>

<Title title="Search IP Address" />
<form on:submit|preventDefault={getIPDetail} class="p-4 text-left" autocomplete="off">
  <div class="relative flex flex-wrap items-stretch rounded w-full">
    <input
      bind:this={input}
      type="text"
      name="search"
      bind:value={ip}
      class="transition duration-150 ease-in-out leading-tight flex flex-auto p-4 rounded border-2 focus:outline-none focus:ring-0
        bg-white border-slate-400 text-gray-600 focus:border-slate-800
        dark:text-gray-200 dark:bg-slate-700 dark:border-slate-600 dark:focus:border-slate-200
        "
    />
    <button
      type="submit"
      class="absolute top-0 right-0 h-full rounded-r px-4 focus:outline-none border-transparent  border border-l-0
       text-gray-600 dark:text-gray-200"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
        />
      </svg>
    </button>
  </div>
  <span class="text-xs text-left" class:text-red-400={error}>
    {error ? error : "please insert ip address and hit enter"}
  </span>
</form>
<hr class="border-slate-100 dark:border-slate-900" />
<div class="relative p-4" class:animate-pulse={loading}>
  <Detail {detail} />
</div>
