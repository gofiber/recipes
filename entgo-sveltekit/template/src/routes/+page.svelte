<script lang="ts">
  import { todos, add, loading, remove, update, type TodoType } from "$lib/store/todos";

  let newTodo = "";

  async function handleAdd() {
    if (newTodo) {
      await add({ content: newTodo });
      newTodo = "";
    }
  }
  async function toggleTodo(todo: TodoType) {
    const updatedTodo = { ...todo, completed: !todo.completed };
    await update(todo.id, updatedTodo);
  }
</script>

<svelte:head>
  <title>Todo App</title>
</svelte:head>
<div class="min-h-screen flex flex-col">
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-3xl font-semibold mb-4">Todo List</h1>
    <div class="flex mb-4">
      <input
        class="flex-grow mr-4 py-2 px-3 border border-gray-300 rounded"
        type="text"
        placeholder="Add Todo"
        bind:value={newTodo}
      />
      <button
        class="px-4 py-2 bg-teal-500 text-white rounded hover:bg-teal-600"
        on:click={handleAdd}
      >
        Add
      </button>
    </div>
    <div>
      {#if $todos.length === 0}
        <p class="text-gray-500">No todos yet.</p>
      {/if}
      {#each $todos as todo}
        <div class="flex items-center mb-2">
          <input
            class="mr-2"
            type="checkbox"
            bind:checked={todo.completed}
            disabled={$loading}
            on:click={() => toggleTodo(todo)}
          />
          <p class="flex-grow" class:selected={todo.completed} class:text-gray-500={todo.completed}>
            {todo.content}
          </p>
          <button
            class="px-2 py-1 text-red-500 hover:text-red-700"
            disabled={$loading}
            on:click={() => remove(todo.id)}
          >
            Remove
          </button>
        </div>
      {/each}
    </div>
  </div>
</div>
