import { cn } from "@/lib/utils";

const SPACE_SIZES: Record<string, string> = {
  xxsmall: "my-0.5",
  xsmall: "my-1",
  small: "my-2",
  medium: "my-4",
  large: "my-8",
  xlarge: "my-12",
  xxlarge: "my-16",
  xxxlarge: "my-24",
};

interface ISpacerProps {
  className?: string;
  size: keyof typeof SPACE_SIZES;
}

const Spacer = ({ className, size }: ISpacerProps) => {
  return <div className={cn(SPACE_SIZES[size], className)} />;
};

export default Spacer;
