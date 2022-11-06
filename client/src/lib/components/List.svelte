<script lang="ts">
    import type { Item } from "$lib/types";
    export let items: Item[]; 
    export let category: string = '';
</script>

<section>

    <div class="task-wrapper" id={category}>
        {#each items as item(item.id)}
            <div class="m-3 p-3 bg-slate-100">
                <p>{item.name}</p>
            </div>
        {/each}
    </div>
</section>