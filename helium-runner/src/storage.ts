import * as events from './events';

interface BuildStorage {
    create(e: events.BrigadeEvent, project: events.Project, size?: string): Promise<string>
    destroy(): Promise<boolean>
}
