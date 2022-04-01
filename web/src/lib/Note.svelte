<script lang="ts">
    export let id: number
    export let title: string
    export let content: string
    export let tags: string[]
</script>

<div class="card w-96 bg-base-100 shadow-xl m-4">
  <figure><img src="https://api.lorem.space/image/shoes?w=400&h=225" alt="Shoes" /></figure>
  <div class="card-body">

    <h2 class="card-title text-2xl uppercase">
        {title}
    </h2>

    <p>{content}</p>

    <div class="card-actions justify-end">
      {#each tags as tag (tag)}
        <div class="badge badge-outline">{tag}</div>
      {/each}
    </div>
  </div>
</div>


<style>
    .badge, .card-title {
        color: #ff3e00;
    }
</style>