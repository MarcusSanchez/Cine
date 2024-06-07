import mascot from "@/../public/mascot.png";

export default function NotFound() {
  return (
    <div className="container max-w-[1200px] mb-8">
      <div className="flex flex-col md:grid md:grid-cols-5 ">
        <div className="order-last text-center md:order-first col-span-3 flex flex-col justify-center">
          <h1 className="text-4xl md:text-6xl lg:text-8xl font-bold text-brand-yellow">Uh oh. 404!</h1>
          <p className="text-sm sm:text-base md:text-xl lg:text-2xl text-brand-light">
            Not the "Godzilla" you were looking for? Looks like the page you visited went extinct...
          </p>
        </div>
        <div className="w-full flex justify-center col-span-2">
          <img src={mascot.src} alt="cinema-mascot" className="max-h-[10rem] md:max-h-[20rem] lg:max-h-[24rem]" />
        </div>
      </div>
    </div>
  );
}