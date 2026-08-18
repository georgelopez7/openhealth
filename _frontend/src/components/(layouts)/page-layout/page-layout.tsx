import Navbar from "#/components/(layouts)/navbar/navbar";
import Spacer from "#/components/(layouts)/spacer/spacer";

interface IPageLayoutProps {
  children: React.ReactNode;
  navbar?: boolean;
}

export const PageLayout = ({ children, navbar = true }: IPageLayoutProps) => {
  return (
    <div className="relative mx-auto flex min-h-screen flex-col overflow-x-hidden py-8 px-4 md:px-8">
      <Spacer size="small" />
      <div className="relative z-10">
        {navbar && (
          <Navbar
            githubLink={import.meta.env.VITE_GITHUB_URL}
            openfgaLink={import.meta.env.VITE_OPENFGA_URL}
          />
        )}
      </div>
      <Spacer size="small" />

      <div className="absolute top-4 left-4 h-4 w-4 pointer-events-none md:h-8 md:w-8">
        <div className="absolute top-0 left-0 h-1 w-full bg-white md:h-2" />
        <div className="absolute top-0 left-0 h-full w-1 bg-white md:w-2" />
      </div>

      <div className="absolute top-4 right-4 h-4 w-4 pointer-events-none md:h-8 md:w-8">
        <div className="absolute top-0 right-0 h-1 w-full bg-white md:h-2" />
        <div className="absolute top-0 right-0 h-full w-1 bg-white md:w-2" />
      </div>

      <div className="absolute bottom-4 left-4 h-4 w-4 pointer-events-none md:h-8 md:w-8">
        <div className="absolute bottom-0 left-0 h-1 w-full bg-white md:h-2" />
        <div className="absolute bottom-0 left-0 h-full w-1 bg-white md:w-2" />
      </div>

      <div className="absolute bottom-4 right-4 h-4 w-4 pointer-events-none md:h-8 md:w-8">
        <div className="absolute bottom-0 right-0 h-1 w-full bg-white md:h-2" />
        <div className="absolute bottom-0 right-0 h-full w-1 bg-white md:w-2" />
      </div>

      <main className="mx-auto flex w-full flex-1 flex-col md:w-[84vw]">
        {children}
      </main>
    </div>
  );
};
