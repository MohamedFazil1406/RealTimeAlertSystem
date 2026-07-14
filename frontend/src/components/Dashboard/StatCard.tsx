type Props = {
  title: string;
  value: number;
};

export default function StatCard({ title, value }: Props) {
  return (
    <div className="bg-white rounded-lg shadow p-6">
      <h2 className="text-gray-500">{title}</h2>

      <p className="text-4xl font-bold mt-2">{value}</p>
    </div>
  );
}
