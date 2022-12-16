<script lang="ts">
    import TaskForm from '$lib/components/TaskForm.svelte';

    import List from '$lib/components/List.svelte';

    let items1 = [
        {id: 1, name: 'item 1'},
        {id: 2, name: 'item 2'},
        {id: 3, name: 'item 3'},
        {id: 4, name: 'item 4'},
        {id: 5, name: 'item 5'},
        {id: 6, name: 'item 6'},
    ];

</script>


<div class="w-screen h-screen">
    <div class="container m-auto">
        <h1 class="text-3xl font-semibold text-zinc-100 p-6">
            Your tasks
        </h1>
        <div class="grid grid-cols-3 gap-4">
                <List items={items1} category="Planning"/>
                <List items={items1} category="In progress"/>
                <List items={items1} category="Finished"/>
        </div>
    </div>

    <TaskForm />

</div> 