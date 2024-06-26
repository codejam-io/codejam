ALTER TABLE statuses ADD COLUMN IF NOT EXISTS sort int;

INSERT INTO statuses (code, title, description) VALUES ('SIGNUP', 'Signups Accepted', 'Users may begin signing up and creating teams.');
INSERT INTO statuses (code, title, description) VALUES ('COMPLETED', 'Complete', 'Competition is over.  Results may still be viewed, but not more voting is allowed');

UPDATE statuses SET sort = 1 WHERE code='PLANNING';
UPDATE statuses SET sort = 2 WHERE code='PUBLISHED';
UPDATE statuses SET sort = 3 WHERE code='SIGNUP';
UPDATE statuses SET sort = 4 WHERE code='STARTED';
UPDATE statuses SET sort = 5 WHERE code='ENDED';
UPDATE statuses SET sort = 6 WHERE code='VOTING';
UPDATE statuses SET sort = 7 WHERE code='COMPLETED';
