import { error,loading,todos, type TodoType } from '$lib/store/todos';
import type { PageLoad } from './$types';

export const load = (async ({fetch}) => {
    loading.set(true); // Set the loading state to true before fetching data
    try {
        const response = await fetch("/api/v1/todo/list");
        const result = await response.json() as TodoType[];
        todos.set(result); // Set the retrieved data to the todos store
    } catch (e) {
        error.set(true); // Set the error state to true if an error occurs
    } finally {
        loading.set(false); // Set the loading state back to false after the fetch operation
    }
    return {};
}) satisfies PageLoad;