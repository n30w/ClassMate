import "@testing-library/jest-dom";
import validatePassword from "@/lib/helpers/passwordValidator";

describe("validatePassword", () => {
  test("returns true for a valid password", () => {
    const validPassword = "!aklklaskdlALSKFJ399davklmasd";
    expect(validatePassword(validPassword)).toBe(true);
  });

  test("returns false for a password without a letter", () => {
    const invalidPassword = "1234567890!@#";
    expect(validatePassword(invalidPassword)).toBe(false);
  });

  test("returns false for a password without a number", () => {
    const invalidPassword = "AbCdEfGhIjKl!@#";
    expect(validatePassword(invalidPassword)).toBe(false);
  });

  test("returns false for a password without a special character", () => {
    const invalidPassword = "AbCdEfGhIjKl1234567890";
    expect(validatePassword(invalidPassword)).toBe(false);
  });

  test("returns false for a password shorter than 8 characters", () => {
    const invalidPassword = "Ab1!";
    expect(validatePassword(invalidPassword)).toBe(false);
  });

  test("returns false for an empty password", () => {
    const emptyPassword = "";
    expect(validatePassword(emptyPassword)).toBe(false);
  });
});
