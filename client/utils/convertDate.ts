export default function convertDate(date: string): string {
  return (
    new Date(date).getUTCDate() +
    "/" +
    (new Date(date).getUTCMonth() + 1) +
    "/" +
    new Date(date).getUTCFullYear()
  );
}
