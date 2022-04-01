<script lang="ts">
    import Note from "./Note.svelte";

    let notes = [
        {
            id: 0,
            title: "Hi there",
            content: "This is a first Go / Svelte app ! I wanna test some new technologies :p",
            tags: ["Test", "Go", "Svelte"]
        },
        {
            id: 1,
            title: "Front End",
            content: "bla blabla bla bla blabla blabla bla bla bla",
            tags: ["TypeScript", "Svelte", "Tailwind", "Vite"]
        },
        {
            id: 2,
            title: "Back End",
            content: "bla blabla bla bla blabla blabla bla bla bla",
            tags: ["Go", "PostGreSQL"]
        },
    ]
</script>


<section class="flex flex-wrap justify-evenly">
    {#each notes as note (note)}
        <Note
            id={note.id}
            title={note.title}
            content={note.content}
            tags={note.tags}
        />
    {/each}
</section>