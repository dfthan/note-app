<script lang="ts">
    export let task:string = "";
    const handleSubmit = () => {
        event?.preventDefault();
        console.log("hello");
    }
</script>

<div>
    <form class="content border-b border-sky-300/25 w-1/4" on:submit={handleSubmit}>
        <input type="text" class="bg-transparent border-none w-max p-4 focus:outline-none" bind:value={task} placeholder="Type task" />
        <button class="bg-green-400/50 border-teal-500 rounded-md p-2 m-1 hover:bg-green-500" >Add task</button>
    </form>
</div>