export function isNullOrEmpty(value?: string | null): boolean {
    return value == null || value.trim() === "";
}