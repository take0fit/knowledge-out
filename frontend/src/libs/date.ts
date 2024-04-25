import { format, parseISO } from 'date-fns';
import { fromZonedTime } from 'date-fns-tz';

// https://qiita.com/suin/items/296740d22624b530f93a
export function isoStringToJstDate(isoString: string): string {
    const utcDate = parseISO(isoString);
    const jstDate = fromZonedTime(utcDate, 'Asia/Tokyo');
    return format(jstDate, 'yyyy-MM-dd');
}