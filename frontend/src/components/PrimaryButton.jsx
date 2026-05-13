function PrimaryButton({ children, className = "", ...props }) {
  return (
    <button
      className={`rounded-md border border-zinc-700 bg-[#171717] px-4 py-2.5 text-sm font-medium text-zinc-50 transition duration-200 hover:border-zinc-600 hover:bg-[#1f1f1f] focus:outline-none focus:ring-2 focus:ring-[#0070f3]/35 disabled:opacity-50 ${className}`}
      {...props}
    >
      {children}
    </button>
  );
}

export default PrimaryButton;
