import "@testing-library/jest-dom";
import passwordValidator from "@/lib/helpers/passwordValidator";

test("valid password", () => {
  expect(passwordValidator("!aklklaskdlALSKFJ399davklmasd.aa")).toBe(true);
});
