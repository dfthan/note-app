<script lang="ts">
    import List from '$lib/components/List.svelte';

    let items1 = [
        {id: 1, name: 'item 1'},
        {id: 2, name: 'item 2'},
        {id: 3, name: 'item 3'},
    ];

</script>


<div class="w-screen h-screen">
    <div class="flex justify-center p-6">
        <h1 class="text-3xl font-bold text-zinc-100">
            Your tasks
        </h1>
    </div>

    <div class="container mx-auto">
        <div class="grid grid-cols-3 gap-4">
                <List items={items1} category="planning"/>
                <List items={items1} category="doing"/>
                <List items={items1} category="finished"/>
        </div>
    </div>

</div> 