<script>
  import { page } from "$app/stores";
</script>

<div class="flex items-stretch w-full rounded-b text-sm">
  <a
    class="flex w-full items-center justify-center h-12 rounded-bl bg-primary 
    {$page.path === '/' ? 'bg-opacity-0' : 'bg-opacity-5'}"
    href="/"
  >
    Home
  </a>
  <a
    class="flex w-full items-center justify-center h-12 rounded-br bg-primary 
    {$page.path === '/detail' ? 'bg-opacity-0' : 'bg-opacity-5'}"
    href="/detail"
  >
    Search IP Address
  </a>
</div>
