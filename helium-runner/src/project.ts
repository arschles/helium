import * as events from './events';

/**
 * ProjectLoader describes a function able to load a Project.
 */
export interface ProjectLoader {
    (projectID: string, projectNS: string): Promise<events.Project>
}

export const loadEmptyProject = (projID: string, projNS: string) => {
    return new Promise<events.Project>((resolve, reject) => {
        let proj = new events.Project();
        resolve(proj);
    });
}
