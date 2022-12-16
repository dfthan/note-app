<script lang="ts">
    import type { Item } from "$lib/types";
    export let items: Item[]; 
    export let category: string = '';
</script>

<section>

    <div class="task-wrapper" id={category}>
        <div class="flex justify-center font-semibold">
            <h1 >{category}</h1>
        </div>
        {#each items as item(item.id)}
            <div class="m-3 p-3 bg-blue-100 rounded" draggable="true">
                <p>{item.name}</p>
            </div>
        {/each}
    </div>
</section>