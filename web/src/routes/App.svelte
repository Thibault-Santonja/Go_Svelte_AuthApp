<div>%svelte.body%</div>
 <script>
   // App.svelte
   import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query'
   import Example from "./Example.svelte";
   const queryClient = new QueryClient()
 </script>

 <QueryClientProvider client={queryClient}>
   <Example />
 </QueryClientProvider>