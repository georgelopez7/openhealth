import { Link } from "@tanstack/react-router";

import GitHubIcon from "@/components/(icons)/github-icon";
import OpenFGAIcon from "@/components/(icons)/openfga-icon";
import OpenhealthIcon from "@/components/(icons)/openhealth-icon";
import { cn } from "@/lib/utils";

interface IProps {
  className?: string;
  githubLink?: string;
  openfgaLink?: string;
}

const Navbar = ({
  className,
  githubLink = "https://github.com",
  openfgaLink = "https://openfga.dev",
}: IProps) => {
  return (
    <nav className={cn("w-full", className)}>
      <div
        className={cn(
          "mx-auto flex w-full items-center justify-between rounded-lg border border-white/20 bg-white/5 px-4 py-2.5 md:w-[84vw]",
        )}
      >
        <Link
          to="/"
          className="flex items-center gap-2 text-base font-semibold tracking-tight text-white transition-colors hover:underline md:text-lg"
        >
          <OpenhealthIcon className="size-7" />
          OpenHealth
        </Link>
        <div className="flex items-center gap-3">
          <a
            href={openfgaLink}
            target="_blank"
            rel="noopener noreferrer"
            className="text-white transition-colors hover:text-white/80"
            aria-label="OpenFGA"
          >
            <OpenFGAIcon className="size-7" />
          </a>
          <a
            href={githubLink}
            target="_blank"
            rel="noopener noreferrer"
            className="text-white transition-colors hover:text-white/80"
            aria-label="GitHub"
          >
            <GitHubIcon size={20} />
          </a>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
