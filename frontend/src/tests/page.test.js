import React from 'react';
import { render, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import Page from "frontend/src/app/(auth)/signup/page";

describe('Page component', () => {
  test('renders correctly', () => {
    const { getByText, getByPlaceholderText } = render(<Page />);
    
    // Assert that necessary elements are present
    expect(getByText('Sign up')).toBeInTheDocument();
    expect(getByPlaceholderText('abc123')).toBeInTheDocument();
    expect(getByPlaceholderText('abc123@nyu.edu')).toBeInTheDocument();
    expect(getByPlaceholderText('••••••••••')).toBeInTheDocument();
    expect(getByPlaceholderText('Re-enter Password')).toBeInTheDocument();
    expect(getByText('Already have an account?')).toBeInTheDocument();
    expect(getByText('Log in')).toBeInTheDocument();
  });

  test('validates password correctly', async () => {
    const { getByText, getByPlaceholderText } = render(<Page />);
    
    // Enter invalid password
    fireEvent.change(getByPlaceholderText('••••••••••'), { target: { value: 'weakpassword' } });
    fireEvent.change(getByPlaceholderText('Re-enter Password'), { target: { value: 'weakpassword' } });
    fireEvent.click(getByText('Sign up'));

    // Assert that password error is displayed
    await waitFor(() => expect(getByText('Password must have at least one letter, one number, one special character, and at least 8 characters long.')).toBeInTheDocument());
  });

  test('validates re-entered password correctly', async () => {
    const { getByText, getByPlaceholderText } = render(<Page />);
    
    // Enter valid password but different re-entered password
    fireEvent.change(getByPlaceholderText('••••••••••'), { target: { value: 'StrongPassword123!' } });
    fireEvent.change(getByPlaceholderText('Re-enter Password'), { target: { value: 'DifferentPassword123!' } });
    fireEvent.click(getByText('Sign up'));

    // Assert that passwords do not match error is displayed
    await waitFor(() => expect(getByText('Passwords do not match.')).toBeInTheDocument());
  });

  // Add more test cases to cover other functionality if needed
});
