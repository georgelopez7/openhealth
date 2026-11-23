import Navbar from "#/components/(layouts)/navbar/navbar";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { EXTERNAL_LINKS } from "#/domain/config";

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
            githubLink={EXTERNAL_LINKS["openhealth.github"]}
            openfgaLink={EXTERNAL_LINKS["openfga.website"]}
          />
        )}
      </div>
      <Spacer size="small" />

      <main className="mx-5 flex w-[calc(100%-2.5rem)] flex-1 flex-col md:mx-auto md:w-[84vw]">
        {children}
      </main>
      <Spacer size="small" />
    </div>
  );
};
