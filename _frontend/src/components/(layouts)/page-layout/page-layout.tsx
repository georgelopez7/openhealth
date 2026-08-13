import Navbar from "#/components/(layouts)/navbar/navbar";
import Spacer from "#/components/(layouts)/spacer/spacer";

interface IPageLayoutProps {
  children: React.ReactNode;
  navbar?: boolean;
}

export const PageLayout = ({ children, navbar = true }: IPageLayoutProps) => {
  return (
    <div className="relative mx-auto flex min-h-screen flex-col p-8">
      <Spacer size="small" />
      <div className="relative z-10">
        {navbar && <Navbar githubLink={import.meta.env.VITE_GITHUB_URL} />}
      </div>
      <Spacer size="xsmall" />

      <div className="absolute top-4 left-4 h-8 w-8 pointer-events-none">
        <div className="absolute top-0 left-0 h-2 w-full bg-white" />
        <div className="absolute top-0 left-0 h-full w-2 bg-white" />
      </div>

      <div className="absolute top-4 right-4 h-8 w-8 pointer-events-none">
        <div className="absolute top-0 right-0 h-2 w-full bg-white" />
        <div className="absolute top-0 right-0 h-full w-2 bg-white" />
      </div>

      <div className="absolute bottom-4 left-4 h-8 w-8 pointer-events-none">
        <div className="absolute bottom-0 left-0 h-2 w-full bg-white" />
        <div className="absolute bottom-0 left-0 h-full w-2 bg-white" />
      </div>

      <div className="absolute bottom-4 right-4 h-8 w-8 pointer-events-none">
        <div className="absolute bottom-0 right-0 h-2 w-full bg-white" />
        <div className="absolute bottom-0 right-0 h-full w-2 bg-white" />
      </div>

      <main className="mx-auto flex w-full max-w-6xl flex-1 flex-col">
        {children}
      </main>
    </div>
  );
};
