import React, { FC } from 'react';
import { Accordion, Badge, Button, Stack } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';
import { useNavigate, useMatch } from 'react-router-dom';

import { useAccordionButton } from 'react-bootstrap/AccordionButton';

import { Icon } from '@/components';

function MenuNode({ menu, callback, activeKey, isLeaf = false }) {
  const { t } = useTranslation('translation', { keyPrefix: 'nav_menus' });
  const accordionClick = useAccordionButton(menu.name);
  const menuOnClick = (evt) => {
    evt.preventDefault();
    evt.stopPropagation();
    if (!isLeaf) {
      accordionClick(evt);
    }
    if (typeof callback === 'function') {
      callback(menu);
    }
  };

  let menuCls = 'text-start text-dark text-nowrap shadow-none bg-body border-0';
  let menuVariant = 'light';
  if (activeKey === menu.name) {
    menuCls = 'text-start text-white text-nowrap shadow-none';
    menuVariant = 'primary';
  }
  return (
    <Button variant={menuVariant} className={menuCls} onClick={menuOnClick}>
      <Stack direction="horizontal">
        {!isLeaf ? <Icon name="chevron-right" className="me-1" /> : null}
        {t(menu.name)}
        {menu.badgeContent ? (
          <Badge bg="dark" className="ms-auto top-0">
            {menu.badgeContent}
          </Badge>
        ) : null}
      </Stack>
    </Button>
  );
}

interface AccordionProps {
  menus: any[];
  path?: string;
}
const AccordionNav: FC<AccordionProps> = ({ menus, path = '/' }) => {
  const navigate = useNavigate();
  const pathMatch = useMatch(`${path}*`);
  if (!menus.length) {
    return null;
  }
  // auto set menu fields
  menus.forEach((m) => {
    if (!Array.isArray(m.children)) {
      m.children = [];
    }
    m.children.forEach((sm) => {
      if (!Array.isArray(sm.children)) {
        sm.children = [];
      }
    });
  });
  const splat = pathMatch && pathMatch.params['*'];
  let activeKey = menus[0].name;
  if (splat) {
    activeKey = splat;
  }
  const menuClick = (clickedMenu) => {
    const menuKey = clickedMenu.name;
    if (clickedMenu.children.length) {
      return;
    }
    if (activeKey !== menuKey) {
      const routePath = `${path}${menuKey}`;
      navigate(routePath);
    }
  };

  let defaultOpenKey;
  menus.forEach((li) => {
    if (li.children.length) {
      const matchedChild = li.children.find((el) => {
        return el.name === activeKey;
      });
      if (matchedChild) {
        defaultOpenKey = li.name;
      }
    }
  });

  return (
    <Accordion defaultActiveKey={defaultOpenKey} flush>
      <Stack direction="vertical" gap={1}>
        {menus.map((li) => {
          return (
            <React.Fragment key={li.name}>
              <MenuNode menu={li} callback={menuClick} activeKey={activeKey} />
              {li.children.length ? (
                <Accordion.Collapse eventKey={li.name} className="ms-4">
                  <Stack direction="vertical" gap={1}>
                    {li.children.map((leaf) => {
                      return (
                        <MenuNode
                          menu={leaf}
                          callback={menuClick}
                          activeKey={activeKey}
                          isLeaf
                          key={leaf.name}
                        />
                      );
                    })}
                  </Stack>
                </Accordion.Collapse>
              ) : null}
            </React.Fragment>
          );
        })}
      </Stack>
    </Accordion>
  );
};

export default AccordionNav;
