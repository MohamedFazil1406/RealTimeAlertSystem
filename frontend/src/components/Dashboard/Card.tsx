type Props = {
  title: string;
  value: number;
};

export default function Card({ title, value }: Props) {
  return (
    <div className="rounded-lg bg-white shadow p-6">
      <h3 className="text-gray-500">{title}</h3>

      <p className="text-4xl font-bold mt-3">{value}</p>
    </div>
  );
}
