import Image from "next/image";

const Discussions = () => {
  return (
    <div className="w-60 h-70 px-4">
      <h1 className="font-bold text-xl border-b-2 border-black mb-4 pb-4">
        Discussions
      </h1>
      <h2 className="text-m font-bold mb-1">
        Balancing Innovation and Maintenance
      </h2>
      <div>
        <Image src={""} alt={""} />
        <h3 className="text-sm text-gray-400 mb-2">
          Neo Alabastro posted on Feb 9, 2024 9:16 PM
        </h3>
      </div>
      <p className="text-sm font-light">In my experience, ...</p>
    </div>
  );
};

export default Discussions;
