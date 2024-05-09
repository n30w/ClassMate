"use server";

export default async function addStudentToClass(courseId: string) {
  return async (formData: FormData) => {
    try {
      const res: Response = await fetch(
        "http://localhost:6789/v1/course/addstudent",
        {
          method: "POST",
          body: JSON.stringify({
            netid: formData.get("netid") as string,
            courseid: courseId,
          }),
        }
      );
      if (res.ok) {
        const response = await res.json();
        console.log(response);
      } else {
        console.error("Failed to add student to the course:", res.statusText);
      }
    } catch (error) {
      console.error("Error adding student to the course:", error);
    }
    return "form submission success";
  };
}
