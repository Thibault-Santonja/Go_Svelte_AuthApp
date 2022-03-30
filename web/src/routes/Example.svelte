<script>
   // Example.svelte
   import { useQuery } from '@sveltestack/svelte-query'
       const queryResult = useQuery('repoData', () =>
         fetch('https://api.github.com/repos/SvelteStack/svelte-query').then(res =>
           res.json()
         )
       )
</script>

 {#if $queryResult.isLoading}
   <span>Loading...</span>
 {:else if $queryResult.error}
   <span>An error has occurred: {$queryResult.error.message}</span>
 {:else}
   <div>
     <h1>{$queryResult.data.name}</h1>
     <p>{$queryResult.data.description}</p>
     <strong>👀 {$queryResult.data.subscribers_count}</strong>{' '}
     <strong>✨ {$queryResult.data.stargazers_count}</strong>{' '}
     <strong>🍴 {$queryResult.data.forks_count}</strong>
   </div>
 {/if}