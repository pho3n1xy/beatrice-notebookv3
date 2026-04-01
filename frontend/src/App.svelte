<script>
  import { onMount} from 'svelte';
  import {SaveDream, GetDreamsDesc} from '../wailsjs/go/main/App.js'


  let dreams = [];
  let currentDream = { title: '', content: '', date: new Date().toISOString().split('T')[0], moonphase: ''};

  onMount(async () =>{
    try {
      dreams = await GetDreamsDesc();
    } catch (err){
      console.error("Failed to load dreams:", err)
    }
  });

  async function saveDream(){
    if (!currentDream.content){
       alert("Please write something in your dream first!");
       return;
    }

    try{
      // Pass the data to your Go backend!
      const newDream = await SaveDream(currentDream.title, currentDream.content, currentDream.date, currentDream.moonphase);

      // Add the newly saved dream (which now has an ID from SQLite) to the top of our list
      dreams = [newDream, ...dreams];

      // Reset the form
      currentDream = {title: '', content: '', date: new Date().toISOString().split('T')[0], moonphase: ''};
    }catch(err){
     console.error("Failed to save dream:", err);
    }
  }
</script>

<main class="flex h-screen bg-gray-100 font-sans text-gray-900">
  <!-- Sidebar -->
  <aside class="w-64 bg-white border-r border-gray-200 flex flex-col">
    <div class="p-4 border-b border-gray-200">
      <h1 class="text-xl font-bold text-indigo-600">Dream Journal</h1>
    </div>
    <div class="flex-1 overflow-y-auto p-4 space-y-2">
      {#if dreams.length === 0}
        <p class="text-gray-500 italic text-sm">No dreams recorded yet.</p>
      {:else}
        {#each dreams as dream}
          <button class="w-full text-left p-2 rounded hover:bg-indigo-50 hover:text-indigo-700 transition-colors">
            <div class="font-medium truncate">{dream.title || "Untitled Dream"}</div>
            <div class="text-xs text-gray-500">{dream.date}</div>
          </button>
        {/each}
      {/if}
    </div>
    <div class="p-4 border-t border-gray-200">
      <button
        class="w-full bg-indigo-50 text-indigo-700 py-2 rounded font-medium hover:bg-indigo-100 transition-colors"
        on:click={() => currentDream = { title: '', content: '', date: new Date().toISOString().split('T')[0] }}
      >
        + New Dream
      </button>
    </div>
  </aside>

  <!-- Main Editor Area -->
  <section class="flex-1 flex flex-col p-8 bg-gray-50">
    <div class="max-w-3xl mx-auto w-full flex-1 flex flex-col bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">

      <!-- Editor Header -->
      <div class="p-6 border-b border-gray-100 flex gap-4">
        <input
          type="date"
          bind:value={currentDream.date}
          class="px-3 py-2 bg-gray-50 border border-gray-200 rounded focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent text-gray-700"
        />
        <input
          type="text"
          placeholder="Title your dream..."
          bind:value={currentDream.title}
          class="flex-1 px-3 py-2 text-lg font-medium bg-transparent border-none focus:outline-none placeholder-gray-400"
        />
      </div>

      <!-- Editor Body -->
      <textarea
        placeholder="Write down what you remember..."
        bind:value={currentDream.content}
        class="flex-1 p-6 resize-none bg-transparent border-none focus:outline-none text-gray-700 leading-relaxed"
      ></textarea>

      <!-- Editor Footer -->
      <div class="p-4 bg-gray-50 border-t border-gray-100 flex justify-end">
        <button
          on:click={saveDream}
          class="px-6 py-2 bg-indigo-600 text-white font-medium rounded hover:bg-indigo-700 transition-colors shadow-sm"
        >
          Save Dream
        </button>
      </div>

    </div>
  </section>
</main>
