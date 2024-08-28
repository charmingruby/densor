interface CategoryTagProps {
  category: string
}

export function CategoryTag({ category }: CategoryTagProps) {
  return (
    <small className="px-2 py-1 rounded-full bg-primary text-xs text-primary-foreground">
      {category}
    </small>
  )
}
