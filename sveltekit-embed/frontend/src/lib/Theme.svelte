<script>
  let theme = "light";
  try {
    theme = localStorage.theme;
  } catch (e) {}
  const toggle = () => {
    const { classList } = document.querySelector("html");
    classList.remove(theme);
    theme = theme === "light" ? "dark" : "light";
    classList.add(theme);
    try {
      localStorage.theme = theme;
    } catch (e) {}
  };
</script>

<button
  aria-label="Toggle Theme"
  class="flex items-center justify-center rounded-full w-12 h-12"
  on:click={toggle}
>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    class="stroke-current text-gray-600 dark:text-gray-200 w-6 h-6"
  >
    {#if theme === "light"}
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
      />
    {:else}
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
      />
    {/if}
  </svg>
  <div class="sr-only">Toggle Theme</div>
</button>
