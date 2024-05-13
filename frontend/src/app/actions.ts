"use server";

import { redirect } from "next/navigation";

// Adapted from:
// https://nextjs.org/docs/app/building-your-application/data-fetching/server-actions-and-mutations#redirecting
export async function createCourse(id: string) {
  try {
    // ...
  } catch (error) {
    // ...
  }

  //   revalidateTag("posts"); // Update cached posts
  redirect(`/post/${id}`); // Navigate to the new post page
}

export async function navigate(data: FormData) {
  redirect(`/posts/${data.get("id")}`);
}
