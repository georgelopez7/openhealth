import { render } from "@testing-library/react";
import { describe, expect, it } from "vitest";
import Spacer from "./spacer";

describe("Spacer", () => {
  it("should render a div element", () => {
    const { container } = render(<Spacer size="small" />);
    expect(container.firstChild).toBeInstanceOf(HTMLDivElement);
  });

  it("should apply the correct size class", () => {
    const { container } = render(<Spacer size="medium" />);
    expect((container.firstChild as HTMLElement).className).toContain("my-4");
  });

  it("should apply a custom className", () => {
    const { container } = render(
      <Spacer size="large" className="custom-class" />,
    );
    const className = (container.firstChild as HTMLElement).className;
    expect(className).toContain("my-8");
    expect(className).toContain("custom-class");
  });

  it("should render with xxlarge size", () => {
    const { container } = render(<Spacer size="xxlarge" />);
    expect((container.firstChild as HTMLElement).className).toContain("my-16");
  });
});
