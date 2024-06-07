export default function Footer() {
  return (
    <footer className="border-y-2 border-brand-yellow flex flex-wrap gap-2 justify-center mt-[60px] p-5">
      <p className="font-semibold italic text-brand-yellow">
        Made by Marcus Sanchez
      </p>
      <p className="font-semibold italic text-brand-yellow">
        |
      </p>
      <p className="font-semibold italic text-brand-yellow">
        Repository: {" "}
        <a
          href="https://github.com/MarcusSanchez/Cine"
          className="text-brand-light hover:text-white hover:underline"
        >Here</a>
      </p>
    </footer>
  );
}