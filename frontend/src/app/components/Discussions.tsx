import Image from "next/image"

const Discussions = () => {
  return (
    <div class="w-60 h-70 px-4">
        <h1 class="font-bold text-xl border-b-2 border-black mb-4 pb-4">Discussions</h1>
        <h2 class="text-m font-bold mb-1">Balancing Innovation and Maintenance</h2>
        <div>
            <Image />
            <h3 class="text-sm text-gray-400 mb-2">Neo Alabastro posted on Feb 9, 2024 9:16 PM</h3>
        </div>
        <p class="text-sm font-light">In my experience, ...</p>
    </div>
  )
}

export default Discussions