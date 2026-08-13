import { Link } from "@tanstack/react-router";

import GitHubIcon from "@/components/(icons)/github-icon";
import OpenhealthIcon from "@/components/(icons)/openhealth-icon";
import { cn } from "@/lib/utils";

interface IProps {
  className?: string;
  githubLink?: string;
}

const Navbar = ({ className, githubLink = "https://github.com" }: IProps) => {
  return (
    <nav className={cn("w-full", className)}>
      <div
        className={cn(
          "mx-auto flex max-w-6xl items-center justify-between border border-border bg-secondary/80 px-6 py-2.5 rounded-md",
        )}
      >
        <Link
          to="/"
          className="flex items-center gap-2 text-lg font-semibold tracking-tight text-foreground transition-colors hover:underline"
        >
          <OpenhealthIcon className="size-7" />
          OpenHealth
        </Link>
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
    </nav>
  );
};

export default Navbar;
