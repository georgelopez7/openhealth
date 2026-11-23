import { createFileRoute, Link } from "@tanstack/react-router";
import GitHubIcon from "#/components/(icons)/github-icon";
import { PageLayout } from "#/components/(layouts)/page-layout/page-layout";
import Spacer from "#/components/(layouts)/spacer/spacer";
import { EXTERNAL_LINKS } from "#/domain/config";
import { cn } from "#/lib/utils";

export const Route = createFileRoute("/")({ component: Home });

function Home() {
  return (
    <PageLayout navbar={false}>
      <div className="flex flex-1 items-center justify-center">
        <div className="flex w-full bg-[#101010] p-3 sm:p-4">
          <div className="flex w-full flex-col items-center justify-center border border-neutral-700 bg-black px-12 py-14 text-center sm:px-16 sm:py-20">
            <h1 className="text-center text-3xl font-bold sm:text-5xl md:text-6xl">
              OpenHealth
            </h1>
            <Spacer size="small" className="my-5 sm:my-3" />
            <div className="flex gap-2">
              <Link
                to="/dashboard"
                className={cn(
                  "inline-flex h-11 items-center justify-center gap-3 whitespace-nowrap rounded-lg border border-white/20 bg-white/5 px-3 text-xs font-medium text-white transition-colors hover:bg-white/10 sm:px-4 sm:text-base",
                )}
              >
                Get Started
              </Link>
              <a
                href={EXTERNAL_LINKS["openhealth.github"]}
                target="_blank"
                rel="noopener noreferrer"
                className={cn(
                  "inline-flex size-11 items-center justify-center rounded-lg border border-white/20 bg-white/5 text-white transition-colors hover:bg-white/10",
                )}
                aria-label="GitHub"
              >
                <GitHubIcon className="size-5" />
              </a>
            </div>
            <Spacer size="small" />
            <div className="inline-flex items-center gap-2 px-4 py-2 text-[10px] text-white/60">
              <span>
                Powered by{" "}
                <a
                  href={EXTERNAL_LINKS["openfga.website"]}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="font-semibold underline underline-offset-2 transition-colors hover:text-white"
                >
                  OpenFGA
                </a>
              </span>
            </div>
          </div>
        </div>
      </div>
    </PageLayout>
  );
}
