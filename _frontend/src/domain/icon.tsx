import { User } from "lucide-react";
import type { ComponentType } from "react";
import AshIcon from "#/components/(icons)/ash-icon";
import MistyIcon from "#/components/(icons)/misty-icon";
import NurseJoy from "#/components/(icons)/nurse-joy-icon";
import ProfessorOak from "#/components/(icons)/professor-oak-icon";
import BrockIcon from "#/components/(icons)/brock-icon";

export const AccountIconMap: Record<
  string,
  ComponentType<{ className?: string }>
> = {
  "ash-ketchum": AshIcon,
  "misty-waterflower": MistyIcon,
  "nurse-joy": NurseJoy,
  "professor-oak": ProfessorOak,
  "brock-harrison": BrockIcon,
};

export interface AccountIconProps {
  avatar: string;
  className?: string;
}

export const AccountIcon = ({ avatar, className }: AccountIconProps) => {
  const Icon = AccountIconMap[avatar] ?? User;
  return <Icon className={className} />;
};
